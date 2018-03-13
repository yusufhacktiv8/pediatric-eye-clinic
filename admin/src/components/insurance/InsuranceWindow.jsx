import React from 'react';
import { Modal, Form, Input } from 'antd';

const FormItem = Form.Item;

const InsuranceWindow = ({ visible, onCancel, onCreate, form, insurance }) => {
  const { getFieldDecorator } = form;

  return (
    <Modal
      visible={visible}
      title="Insurance"
      okText="Save"
      onCancel={onCancel}
      onOk={onCreate}
    >
      <Form layout="vertical">
        <FormItem label="Code">
          {getFieldDecorator('code', {
            initialValue: insurance.code,
            rules: [
              { required: true, message: 'Please input the code' },
            ],
          })(
            <Input />,
          )}
        </FormItem>
        <FormItem label="Name">
          {getFieldDecorator('name', {
            initialValue: insurance.name,
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

export default Form.create()(InsuranceWindow);
