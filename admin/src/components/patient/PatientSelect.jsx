import React, { Component } from 'react';
import { Select } from 'antd';
import axios from 'axios';
import constant from '../../constant';

const Option = Select.Option;

const PATIENTS_URL = `${constant.serverUrl}/patients_all`;

class PatientSelect extends Component {
  constructor(props) {
    super(props);

    const value = this.props.value || {};
    this.state = {
      patients: [],
      id: value.id,
      loading: false,
    };
  }

  componentDidMount() {
    this.getPatients();
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

  getPatients() {
    this.setState({
      loading: true,
    });
    axios.get(PATIENTS_URL, { params: {} })
      .then((response) => {
        this.setState({
          patients: response.data,
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
        placeholder="Select Patient"
        optionFilterProp="children"
        onChange={(e) => { this.handleChange(e); }}
        filterOption={
          (input, option) => option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0}
      >
        {
          this.state.patients.map(patient => (
            <Option key={patient.id} value={patient.id}>{patient.name}</Option>
          ))
        }
      </Select>
    );
  }
}

export default PatientSelect;
