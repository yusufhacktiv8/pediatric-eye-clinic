import React, { Component } from 'react';
import axios from 'axios';
import { Table, Button, Input, Row, Col, message, Popconfirm } from 'antd';
import constant from '../../constant';
import OccupationWindow from './OccupationWindow';

const OCCUPATIONS_URL = `${constant.serverUrl}/occupations`;
const Column = Table.Column;

class OccupationList extends Component {
  state = {
    searchText: '',
    occupation: {},
    occupations: [],
    loading: false,
    count: 0,
    currentPage: 1,
    pageSize: 10,
    occupationWindowVisible: false,
  }
  componentDidMount() {
    this.getOccupations();
  }

  getOccupations() {
    this.setState({
      loading: true,
    });
    axios.get(OCCUPATIONS_URL, { params: {
      searchText: this.state.searchText,
      start: (this.state.currentPage - 1) * this.state.pageSize,
      count: this.state.pageSize,
    } })
      .then((response) => {
        this.setState({
          occupations: response.data.occupations,
          count: response.data.count,
          loading: false,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  filterOccupations() {
    this.setState({
      currentPage: 1,
    }, () => { this.getOccupations(); });
  }

  saveOccupation(occupation) {
    const hide = message.loading('Action in progress..', 0);
    axios.post(OCCUPATIONS_URL, occupation)
      .then(() => {
        hide();
        this.handleCancel();
        this.getOccupations();
        message.success('Save occupation success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  deleteOccupation(occupation) {
    const hide = message.loading('Action in progress..', 0);
    axios.delete(`${OCCUPATIONS_URL}/${occupation.code}`)
      .then(() => {
        hide();
        this.getOccupations();
        message.success('Delete occupation success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  openEditWindow(record) {
    this.setState({
      occupation: record,
      occupationWindowVisible: true,
    });
  }

  handleCancel() {
    this.setState({
      occupationWindowVisible: false,
    });
    this.occupationWindow.resetFields();
  }

  handleCreate() {
    this.occupationWindow.validateFields((err, values) => {
      if (err) {
        return;
      }

      // console.log('Received values of form: ', values);
      this.saveOccupation(values);
      this.occupationWindow.resetFields();
      this.setState({ occupationWindowVisible: false });
    });
  }

  pageChanged(page) {
    this.setState({
      currentPage: page,
    }, () => { this.getOccupations(); });
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
                onClick={() => this.filterOccupations()}
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
              dataSource={this.state.occupations}
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
                      title={`Are you sure delete occupation ${record.name}`}
                      onConfirm={() => this.deleteOccupation(record)}
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

        <OccupationWindow
          visible={this.state.occupationWindowVisible}
          onCreate={() => this.handleCreate()}
          onCancel={() => this.handleCancel()}
          occupation={this.state.occupation}
          ref={occupationWindow => (this.occupationWindow = occupationWindow)}
        />
      </div>
    );
  }
}

export default OccupationList;
