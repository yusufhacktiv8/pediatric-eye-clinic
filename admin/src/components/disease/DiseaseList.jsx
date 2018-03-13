import React, { Component } from 'react';
import axios from 'axios';
import { Table, Button, Input, Row, Col, message, Popconfirm } from 'antd';
import constant from '../../constant';
import DiseaseWindow from './DiseaseWindow';

const DISEASES_URL = `${constant.serverUrl}/diseases`;
const Column = Table.Column;

class DiseaseList extends Component {
  state = {
    searchText: '',
    disease: {},
    diseases: [],
    loading: false,
    count: 0,
    currentPage: 1,
    pageSize: 10,
    diseaseWindowVisible: false,
  }
  componentDidMount() {
    this.getDiseases();
  }

  getDiseases() {
    this.setState({
      loading: true,
    });
    axios.get(DISEASES_URL, { params: {
      searchText: this.state.searchText,
      start: (this.state.currentPage - 1) * this.state.pageSize,
      count: this.state.pageSize,
    } })
      .then((response) => {
        this.setState({
          diseases: response.data.diseases,
          count: response.data.count,
          loading: false,
        });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  filterDiseases() {
    this.setState({
      currentPage: 1,
    }, () => { this.getDiseases(); });
  }

  saveDisease(disease) {
    const hide = message.loading('Action in progress..', 0);
    axios.post(DISEASES_URL, disease)
      .then(() => {
        hide();
        this.handleCancel();
        this.getDiseases();
        message.success('Save disease success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  deleteDisease(disease) {
    const hide = message.loading('Action in progress..', 0);
    axios.delete(`${DISEASES_URL}/${disease.code}`)
      .then(() => {
        hide();
        this.getDiseases();
        message.success('Delete disease success');
      })
      .catch((error) => {
        hide();
        console.error(error);
      });
  }

  openEditWindow(record) {
    this.setState({
      disease: record,
      diseaseWindowVisible: true,
    });
  }

  handleCancel() {
    this.setState({
      diseaseWindowVisible: false,
    });
    this.diseaseWindow.resetFields();
  }

  handleCreate() {
    this.diseaseWindow.validateFields((err, values) => {
      if (err) {
        return;
      }

      // console.log('Received values of form: ', values);
      this.saveDisease(values);
      this.diseaseWindow.resetFields();
      this.setState({ diseaseWindowVisible: false });
    });
  }

  pageChanged(page) {
    this.setState({
      currentPage: page,
    }, () => { this.getDiseases(); });
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
                onClick={() => this.filterDiseases()}
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
              dataSource={this.state.diseases}
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
                      title={`Are you sure delete disease ${record.name}`}
                      onConfirm={() => this.deleteDisease(record)}
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

        <DiseaseWindow
          visible={this.state.diseaseWindowVisible}
          onCreate={() => this.handleCreate()}
          onCancel={() => this.handleCancel()}
          disease={this.state.disease}
          ref={diseaseWindow => (this.diseaseWindow = diseaseWindow)}
        />
      </div>
    );
  }
}

export default DiseaseList;
