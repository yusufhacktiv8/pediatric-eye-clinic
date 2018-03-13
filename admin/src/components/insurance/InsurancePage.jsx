import React from 'react';
import { Layout } from 'antd';
import InsuranceList from './InsuranceList';

const { Header, Content } = Layout;

export default () => (
  <Layout style={{ height: '100%' }}>
    <Header className="page-header">
      <span>Insurances</span>
    </Header>
    <Content className="page-content">
      <InsuranceList />
    </Content>
  </Layout>
);
