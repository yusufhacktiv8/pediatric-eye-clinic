import React from 'react';
import { Modal, Form, Input } from 'antd';

const FormItem = Form.Item;

const DiseaseWindow = ({ visible, onCancel, onCreate, form, disease }) => {
  const { getFieldDecorator } = form;

  return (
    <Modal
      visible={visible}
      title="Disease"
      okText="Save"
      onCancel={onCancel}
      onOk={onCreate}
    >
      <Form layout="vertical">
        <FormItem label="Code">
          {getFieldDecorator('code', {
            initialValue: disease.code,
            rules: [
              { required: true, message: 'Please input the code' },
            ],
          })(
            <Input />,
          )}
        </FormItem>
        <FormItem label="Name">
          {getFieldDecorator('name', {
            initialValue: disease.name,
            rules: [
              { required: true, message: 'Please input the name' },
            ],
          })(
            <Input />,
          )}
        </FormItem>
      </Form>
    </Modal>
  );
};

export default Form.create()(DiseaseWindow);
