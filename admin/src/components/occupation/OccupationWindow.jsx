import React from 'react';
import { Modal, Form, Input } from 'antd';

const FormItem = Form.Item;

const OccupationWindow = ({ visible, onCancel, onCreate, form, occupation }) => {
  const { getFieldDecorator } = form;

  return (
    <Modal
      visible={visible}
      title="Occupation"
      okText="Save"
      onCancel={onCancel}
      onOk={onCreate}
    >
      <Form layout="vertical">
        <FormItem label="Code">
          {getFieldDecorator('code', {
            initialValue: occupation.code,
            rules: [
              { required: true, message: 'Please input the code' },
            ],
          })(
            <Input />,
          )}
        </FormItem>
        <FormItem label="Name">
          {getFieldDecorator('name', {
            initialValue: occupation.name,
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

export default Form.create()(OccupationWindow);
