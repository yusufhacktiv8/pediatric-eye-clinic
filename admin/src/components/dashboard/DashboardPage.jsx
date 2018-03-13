import React from 'react';
import { Layout, Row, Col, Card } from 'antd';
import Chart1 from './Chart1';
import Chart2 from './Chart2';

const { Header, Content } = Layout;

export default () => (
  <Layout style={{ height: '100%' }}>
    <Header className="page-header">
      <span>Dashboard</span>
    </Header>
    <Content className="page-content">
      <Row gutter={20}>
        <Col span={12}>
          <Card title="Chart 1" extra={<a href="/#">More</a>} type="inner">
            <Chart2 />
          </Card>
        </Col>
        <Col span={12}>
          <Card title="Chart 2" extra={<a href="/#">More</a>} type="inner">
            <Chart1 />
          </Card>
        </Col>
      </Row>
    </Content>
  </Layout>
);
