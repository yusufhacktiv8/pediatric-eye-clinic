import React, { Component } from 'react';
import axios from 'axios';
import { Table, Button, Input, Row, Col, message, Popconfirm } from 'antd';
import constant from '../../constant';
import PatientWindow from './PatientWindow';

const PATIENTS_URL = `${constant.serverUrl}/patients`;
const Column = Table.Column;

class PatientList extends Component {
  state = {
    searchText: '',
    patient: {},
    patients: [],
    loading: false,
    count: 0,
    currentPage: 1,
    pageSize: 10,
    patientWindowVisible: false,
  }
  componentDidMount() {
    this.getPatients();
  }

  getPatients() {
    this.setState({
      loading: true,
    });
    axios.get(PATIENTS_URL, {
      params: {
        searchText: this.state.searchText,
        start: (this.state.currentPage - 1) * this.state.pageSize,
        count: this.state.pageSize,
      },
      headers: {
        Authorization: `Bearer ${window.sessionStorage.getItem('token')}`,
      },
    })
      .then((response) => {
        this.setState({
          patients: response.data.patients,
          count: response.data.count,
          loading: false,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  filterPatients() {
    this.setState({
      currentPage: 1,
    }, () => { this.getPatients(); });
  }

  savePatient(patient) {
    const hide = message.loading('Action in progress..', 0);
    axios.post(PATIENTS_URL, patient)
      .then(() => {
        hide();
        this.handleCancel();
        this.getPatients();
        message.success('Save patient success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  deletePatient(patient) {
    const hide = message.loading('Action in progress..', 0);
    axios.delete(`${PATIENTS_URL}/${patient.code}`)
      .then(() => {
        hide();
        this.getPatients();
        message.success('Delete patient success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  openEditWindow(record) {
    this.setState({
      patient: record,
      patientWindowVisible: true,
    });
  }

  handleCancel() {
    this.setState({
      patientWindowVisible: false,
    });
    this.patientWindow.resetFields();
  }

  handleCreate() {
    this.patientWindow.validateFields((err, values) => {
      if (err) {
        return;
      }

      // console.log('Received values of form: ', values);
      this.savePatient(values);
      // this.patientWindow.resetFields();
      this.setState({ patientWindowVisible: false });
    });
  }

  pageChanged(page) {
    this.setState({
      currentPage: page,
    }, () => { this.getPatients(); });
  }

  render() {
    return (
      <div>
        <Row gutter={10}>
          <Col span={8}>
            <Input
              value={this.state.searchText}
              onChange={(e) => {
                this.setState({
                  searchText: e.target.value,
                });
              }}
              placeholder="Code or name"
            />
          </Col>
          <Col span={16}>
            <span>
              <Button
                shape="circle"
                icon="search"
                onClick={() => this.filterPatients()}
                style={{ marginRight: 15 }}
              />
              <Button
                type="primary"
                shape="circle"
                icon="plus"
                onClick={() => this.openEditWindow({})}
              />
            </span>
          </Col>
        </Row>
        <Row>
          <Col span={24}>
            <Table
              dataSource={this.state.patients}
              style={{ marginTop: 20 }}
              rowKey="id"
              loading={this.state.loading}
              pagination={{
                total: this.state.count,
                current: this.state.currentPage,
                pageSize: this.state.pageSize,
              }}
              onChange={pagination => this.pageChanged(pagination.current)}
              size="small"
            >
              <Column
                title="Code"
                dataIndex="code"
                key="code"
              />
              <Column
                title="Name"
                dataIndex="name"
                key="name"
              />
              <Column
                title="Action"
                key="action"
                render={(text, record) => (
                  <span>
                    <Button
                      icon="ellipsis"
                      size="small"
                      onClick={() => this.openEditWindow(record)}
                      style={{ marginRight: 5 }}
                    />
                    <Popconfirm
                      title={`Are you sure delete patient ${record.name}`}
                      onConfirm={() => this.deletePatient(record)}
                      okText="Yes" cancelText="No"
                    >
                      <Button
                        type="danger"
                        icon="delete"
                        size="small"
                      />
                    </Popconfirm>
                  </span>
                )}
              />
            </Table>
          </Col>
        </Row>

        <PatientWindow
          visible={this.state.patientWindowVisible}
          onCreate={() => this.handleCreate()}
          onCancel={() => this.handleCancel()}
          patient={this.state.patient}
          ref={patientWindow => (this.patientWindow = patientWindow)}
        />
      </div>
    );
  }
}

export default PatientList;
