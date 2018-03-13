import React from 'react';
import { Layout } from 'antd';
import DiseaseList from './DiseaseList';

const { Header, Content } = Layout;

export default () => (
  <Layout style={{ height: '100%' }}>
    <Header className="page-header">
      <span>Diseases</span>
    </Header>
    <Content className="page-content">
      <DiseaseList />
    </Content>
  </Layout>
);
