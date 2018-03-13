import React from 'react';
import { Layout } from 'antd';
import MedicalRecordList from './MedicalRecordList';

const { Header, Content } = Layout;

export default () => (
  <Layout style={{ height: '100%' }}>
    <Header className="page-header">
      <span>MedicalRecords</span>
    </Header>
    <Content className="page-content">
      <MedicalRecordList />
    </Content>
  </Layout>
);
