import React, { Component } from 'react';
import { Select } from 'antd';
import axios from 'axios';
import constant from '../../constant';

const Option = Select.Option;

const OCCUPATIONS_URL = `${constant.serverUrl}/occupations_all`;

class OccupationSelect extends Component {
  constructor(props) {
    super(props);

    const value = this.props.value || {};
    this.state = {
      occupations: [],
      id: value.id,
      loading: false,
    };
  }

  componentDidMount() {
    this.getOccupations();
  }

  componentWillReceiveProps(nextProps) {
    if ('value' in nextProps) {
      const value = nextProps.value;
      if (value) {
        this.setState(value);
      } else {
        this.setState({
          id: undefined,
        });
      }
    }
  }

  getOccupations() {
    this.setState({
      loading: true,
    });
    axios.get(OCCUPATIONS_URL, { params: {} })
      .then((response) => {
        this.setState({
          occupations: response.data,
          loading: false,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  handleChange(value) {
    const temp = {
      id: value,
    };

    this.setState(temp);
    this.triggerChange(temp);
  }

  triggerChange = (changedValue) => {
    const onChange = this.props.onChange;
    if (onChange) {
      onChange(Object.assign({}, changedValue));
    }
  }

  render() {
    return (
      <Select
        value={this.state.id}
        showSearch
        style={{ width: 200 }}
        placeholder="Select Occupation"
        optionFilterProp="children"
        onChange={(e) => { this.handleChange(e); }}
        filterOption={
          (input, option) => option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0}
      >
        {
          this.state.occupations.map(occupation => (
            <Option key={occupation.id} value={occupation.id}>{occupation.name}</Option>
          ))
        }
      </Select>
    );
  }
}

export default OccupationSelect;
