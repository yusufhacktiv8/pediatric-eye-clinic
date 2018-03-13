import React, { Component } from 'react';
import axios from 'axios';
import { Table, Button, Input, Row, Col, message, Popconfirm } from 'antd';
import constant from '../../constant';
import MedicalRecordWindow from './MedicalRecordWindow';

const MEDICAL_RECORDS_URL = `${constant.serverUrl}/medicalrecords`;
const Column = Table.Column;

class MedicalRecordList extends Component {
  state = {
    searchText: '',
    medicalRecord: {},
    medicalRecords: [],
    loading: false,
    count: 0,
    currentPage: 1,
    pageSize: 10,
    medicalRecordWindowVisible: false,
  }
  componentDidMount() {
    this.getMedicalRecords();
  }

  getMedicalRecords() {
    this.setState({
      loading: true,
    });
    axios.get(MEDICAL_RECORDS_URL, { params: {
      searchText: this.state.searchText,
      start: (this.state.currentPage - 1) * this.state.pageSize,
      count: this.state.pageSize,
    } })
      .then((response) => {
        this.setState({
          medicalRecords: response.data.medicalRecords,
          count: response.data.count,
          loading: false,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  filterMedicalRecords() {
    this.setState({
      currentPage: 1,
    }, () => { this.getMedicalRecords(); });
  }

  saveMedicalRecord(medicalRecord) {
    const hide = message.loading('Action in progress..', 0);
    axios.post(MEDICAL_RECORDS_URL, medicalRecord)
      .then(() => {
        hide();
        this.handleCancel();
        this.getMedicalRecords();
        message.success('Save medicalRecord success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  deleteMedicalRecord(medicalRecord) {
    const hide = message.loading('Action in progress..', 0);
    axios.delete(`${MEDICAL_RECORDS_URL}/${medicalRecord.code}`)
      .then(() => {
        hide();
        this.getMedicalRecords();
        message.success('Delete medicalRecord success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  openEditWindow(record) {
    this.setState({
      medicalRecord: record,
      medicalRecordWindowVisible: true,
    });
  }

  handleCancel() {
    this.setState({
      medicalRecordWindowVisible: false,
    });
    this.medicalRecordWindow.resetFields();
  }

  handleCreate() {
    this.medicalRecordWindow.validateFields((err, values) => {
      if (err) {
        return;
      }

      // console.log('Received values of form: ', values);
      this.saveMedicalRecord(values);
      this.medicalRecordWindow.resetFields();
      this.setState({ medicalRecordWindowVisible: false });
    });
  }

  pageChanged(page) {
    this.setState({
      currentPage: page,
    }, () => { this.getMedicalRecords(); });
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
                onClick={() => this.filterMedicalRecords()}
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
              dataSource={this.state.medicalRecords}
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
                title="Patient"
                dataIndex="patient.name"
                key="patient.name"
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
                      title={`Are you sure delete medicalRecord ${record.code}`}
                      onConfirm={() => this.deleteMedicalRecord(record)}
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

        <MedicalRecordWindow
          visible={this.state.medicalRecordWindowVisible}
          onCreate={() => this.handleCreate()}
          onCancel={() => this.handleCancel()}
          medicalRecord={this.state.medicalRecord}
          ref={medicalRecordWindow => (this.medicalRecordWindow = medicalRecordWindow)}
        />
      </div>
    );
  }
}

export default MedicalRecordList;
