import React, { Component } from 'react';
import { Select } from 'antd';
import axios from 'axios';
import constant from '../../constant';

const Option = Select.Option;

const INSURANCES_URL = `${constant.serverUrl}/insurances_all`;

class InsuranceSelect extends Component {
  constructor(props) {
    super(props);

    const value = this.props.value || {};
    this.state = {
      insurances: [],
      id: value.id,
      loading: false,
    };
  }

  componentDidMount() {
    this.getInsurances();
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

  getInsurances() {
    this.setState({
      loading: true,
    });
    axios.get(INSURANCES_URL, { params: {} })
      .then((response) => {
        this.setState({
          insurances: response.data,
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
        placeholder="Select Insurance"
        optionFilterProp="children"
        onChange={(e) => { this.handleChange(e); }}
        filterOption={
          (input, option) => option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0}
      >
        {
          this.state.insurances.map(insurance => (
            <Option key={insurance.id} value={insurance.id}>{insurance.name}</Option>
          ))
        }
      </Select>
    );
  }
}

export default InsuranceSelect;
