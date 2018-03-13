import React from 'react';
import { Modal, Form, Input } from 'antd';

const FormItem = Form.Item;

const UserWindow = ({ visible, onCancel, onCreate, form, user }) => {
  const { getFieldDecorator } = form;

  return (
    <Modal
      visible={visible}
      title="User"
      okText="Save"
      onCancel={onCancel}
      onOk={onCreate}
    >
      <Form layout="vertical">
        <FormItem label="Email">
          {getFieldDecorator('email', {
            initialValue: user.email,
            rules: [
              { required: true, message: 'Please input the email' },
              { type: 'email', message: 'The input is not valid email' },
            ],
          })(
            <Input />,
          )}
        </FormItem>
        <FormItem label="Name">
          {getFieldDecorator('name', {
            initialValue: user.name,
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

export default Form.create()(UserWindow);
