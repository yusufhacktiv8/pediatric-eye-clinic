import React from 'react';
import { Layout } from 'antd';
import OccupationList from './OccupationList';

const { Header, Content } = Layout;

export default () => (
  <Layout style={{ height: '100%' }}>
    <Header className="page-header">
      <span>Occupations</span>
    </Header>
    <Content className="page-content">
      <OccupationList />
    </Content>
  </Layout>
);
