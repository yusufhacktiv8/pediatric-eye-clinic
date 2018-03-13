import React, { Component } from 'react';
import axios from 'axios';
import { Table, Button, Input, Row, Col, message, Popconfirm } from 'antd';
import constant from '../../constant';
import UserWindow from './UserWindow';

const USERS_URL = `${constant.serverUrl}/users`;
const Column = Table.Column;

class UserList extends Component {
  state = {
    searchText: '',
    user: {},
    users: [],
    loading: false,
    count: 0,
    currentPage: 1,
    pageSize: 10,
    userWindowVisible: false,
  }
  componentDidMount() {
    this.getUsers();
  }

  getUsers() {
    this.setState({
      loading: true,
    });
    axios.get(USERS_URL, { params: {
      searchText: this.state.searchText,
      start: (this.state.currentPage - 1) * this.state.pageSize,
      count: this.state.pageSize,
    } })
      .then((response) => {
        this.setState({
          users: response.data.users,
          count: response.data.count,
          loading: false,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  filterUsers() {
    this.setState({
      currentPage: 1,
    }, () => { this.getUsers(); });
  }

  saveUser(user) {
    const hide = message.loading('Action in progress..', 0);
    axios.post(USERS_URL, user)
      .then(() => {
        hide();
        this.handleCancel();
        this.getUsers();
        message.success('Save user success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  deleteUser(user) {
    const hide = message.loading('Action in progress..', 0);
    axios.delete(`${USERS_URL}/${user.email}`)
      .then(() => {
        hide();
        this.getUsers();
        message.success('Delete user success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  openEditWindow(record) {
    this.setState({
      user: record,
      userWindowVisible: true,
    });
  }

  handleCancel() {
    this.setState({
      userWindowVisible: false,
    });
    this.userWindow.resetFields();
  }

  handleCreate() {
    this.userWindow.validateFields((err, values) => {
      if (err) {
        return;
      }

      // console.log('Received values of form: ', values);
      this.saveUser(values);
      this.userWindow.resetFields();
      this.setState({ userWindowVisible: false });
    });
  }

  pageChanged(page) {
    this.setState({
      currentPage: page,
    }, () => { this.getUsers(); });
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
              placeholder="Email or name"
            />
          </Col>
          <Col span={16}>
            <span>
              <Button
                shape="circle"
                icon="search"
                onClick={() => this.filterUsers()}
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
              dataSource={this.state.users}
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
                title="Email"
                dataIndex="email"
                key="email"
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
                      title={`Are you sure delete user ${record.name}`}
                      onConfirm={() => this.deleteUser(record)}
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

        <UserWindow
          visible={this.state.userWindowVisible}
          onCreate={() => this.handleCreate()}
          onCancel={() => this.handleCancel()}
          user={this.state.user}
          ref={userWindow => (this.userWindow = userWindow)}
        />
      </div>
    );
  }
}

export default UserList;
