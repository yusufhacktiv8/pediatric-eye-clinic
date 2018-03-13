import React from 'react';
import { Layout } from 'antd';
import PatientList from './PatientList';

const { Header, Content } = Layout;

export default () => (
  <Layout style={{ height: '100%' }}>
    <Header className="page-header">
      <span>Patients</span>
    </Header>
    <Content className="page-content">
      <PatientList />
    </Content>
  </Layout>
);
