import React, { Component } from 'react';
import axios from 'axios';
import { Table, Button, Input, Row, Col, message, Popconfirm } from 'antd';
import constant from '../../constant';
import InsuranceWindow from './InsuranceWindow';

const INSURANCES_URL = `${constant.serverUrl}/insurances`;
const Column = Table.Column;

class InsuranceList extends Component {
  state = {
    searchText: '',
    insurance: {},
    insurances: [],
    loading: false,
    count: 0,
    currentPage: 1,
    pageSize: 10,
    insuranceWindowVisible: false,
  }
  componentDidMount() {
    this.getInsurances();
  }

  getInsurances() {
    this.setState({
      loading: true,
    });
    axios.get(INSURANCES_URL, { params: {
      searchText: this.state.searchText,
      start: (this.state.currentPage - 1) * this.state.pageSize,
      count: this.state.pageSize,
    } })
      .then((response) => {
        this.setState({
          insurances: response.data.insurances,
          count: response.data.count,
          loading: false,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  filterInsurances() {
    this.setState({
      currentPage: 1,
    }, () => { this.getInsurances(); });
  }

  saveInsurance(insurance) {
    const hide = message.loading('Action in progress..', 0);
    axios.post(INSURANCES_URL, insurance)
      .then(() => {
        hide();
        this.handleCancel();
        this.getInsurances();
        message.success('Save insurance success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  deleteInsurance(insurance) {
    const hide = message.loading('Action in progress..', 0);
    axios.delete(`${INSURANCES_URL}/${insurance.code}`)
      .then(() => {
        hide();
        this.getInsurances();
        message.success('Delete insurance success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  openEditWindow(record) {
    this.setState({
      insurance: record,
      insuranceWindowVisible: true,
    });
  }

  handleCancel() {
    this.setState({
      insuranceWindowVisible: false,
    });
    this.insuranceWindow.resetFields();
  }

  handleCreate() {
    this.insuranceWindow.validateFields((err, values) => {
      if (err) {
        return;
      }

      // console.log('Received values of form: ', values);
      this.saveInsurance(values);
      this.insuranceWindow.resetFields();
      this.setState({ insuranceWindowVisible: false });
    });
  }

  pageChanged(page) {
    this.setState({
      currentPage: page,
    }, () => { this.getInsurances(); });
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
                onClick={() => this.filterInsurances()}
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
              dataSource={this.state.insurances}
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
                      title={`Are you sure delete insurance ${record.name}`}
                      onConfirm={() => this.deleteInsurance(record)}
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

        <InsuranceWindow
          visible={this.state.insuranceWindowVisible}
          onCreate={() => this.handleCreate()}
          onCancel={() => this.handleCancel()}
          insurance={this.state.insurance}
          ref={insuranceWindow => (this.insuranceWindow = insuranceWindow)}
        />
      </div>
    );
  }
}

export default InsuranceList;
