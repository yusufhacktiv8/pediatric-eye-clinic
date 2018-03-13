import React from 'react';
import { Modal, Form, Input } from 'antd';

const FormItem = Form.Item;

const RoleWindow = ({ visible, onCancel, onCreate, form, role }) => {
  const { getFieldDecorator } = form;

  return (
    <Modal
      visible={visible}
      title="Role"
      okText="Save"
      onCancel={onCancel}
      onOk={onCreate}
    >
      <Form layout="vertical">
        <FormItem label="Code">
          {getFieldDecorator('code', {
            initialValue: role.code,
            rules: [
              { required: true, message: 'Please input the code' },
            ],
          })(
            <Input />,
          )}
        </FormItem>
        <FormItem label="Name">
          {getFieldDecorator('name', {
            initialValue: role.name,
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

export default Form.create()(RoleWindow);
