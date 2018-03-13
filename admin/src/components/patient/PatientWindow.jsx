import React from 'react';
import { Modal, Form, Input, DatePicker, Tabs } from 'antd';
import moment from 'moment';
import OccupationSelect from '../occupation/OccupationSelect';
import InsuranceSelect from '../insurance/InsuranceSelect';

const TabPane = Tabs.TabPane;
const FormItem = Form.Item;

const PatientWindow = ({ visible, onCancel, onCreate, form, patient }) => {
  const { getFieldDecorator } = form;
  const fatherOccupation = patient.fatherOccupation && patient.fatherOccupation.id > 0 ?
    patient.fatherOccupation : { id: undefined };
  const motherOccupation = patient.motherOccupation && patient.motherOccupation.id > 0 ?
    patient.motherOccupation : { id: undefined };
  const insurance = patient.insurance && patient.insurance.id > 0 ?
    patient.insurance : { id: undefined };

  return (
    <Modal
      visible={visible}
      title="Patient"
      okText="Save"
      onCancel={onCancel}
      onOk={onCreate}
      bodyStyle={{ paddingTop: 5, height: 407 }}
    >
      <Form layout="vertical">
        <Tabs defaultActiveKey="1">
          <TabPane tab="General" key="1">
            <FormItem label="Code">
              {getFieldDecorator('code', {
                initialValue: patient.code,
                rules: [
                  { required: true, message: 'Please input the code' },
                ],
              })(
                <Input />,
              )}
            </FormItem>
            <FormItem label="Name">
              {getFieldDecorator('name', {
                initialValue: patient.name,
                rules: [
                  { required: true, message: 'Please input the name' },
                ],
              })(
                <Input />,
              )}
            </FormItem>
            <FormItem label="Date of Birth">
              {getFieldDecorator('dateOfBirth', {
                initialValue: moment(patient.dateOfBirth),
                rules: [
                  { required: true, message: 'Please input birth of date' },
                ],
              })(
                <DatePicker />,
              )}
            </FormItem>
            <FormItem label="Address">
              {getFieldDecorator('address', {
                initialValue: patient.address,
              })(
                <Input />,
              )}
            </FormItem>
          </TabPane>
          <TabPane tab="Parents" key="2">
            <FormItem label="Father Name">
              {getFieldDecorator('fatherName', {
                initialValue: patient.fatherName,
              })(
                <Input />,
              )}
            </FormItem>
            <FormItem label="Mother Name">
              {getFieldDecorator('motherName', {
                initialValue: patient.motherName,
              })(
                <Input />,
              )}
            </FormItem>
            <FormItem label="Father Occupation">
              {getFieldDecorator('fatherOccupation', {
                initialValue: fatherOccupation,
              })(
                <OccupationSelect />,
              )}
            </FormItem>
            <FormItem label="Mother Occupation">
              {getFieldDecorator('motherOccupation', {
                initialValue: motherOccupation,
              })(
                <OccupationSelect />,
              )}
            </FormItem>
          </TabPane>
          <TabPane tab="Other" key="3">
            <FormItem label="Insurance">
              {getFieldDecorator('insurance', {
                initialValue: insurance,
              })(
                <InsuranceSelect />,
              )}
            </FormItem>
            <FormItem label="Referral Origin">
              {getFieldDecorator('referralOrigin', {
                initialValue: patient.referralOrigin,
              })(
                <Input />,
              )}
            </FormItem>
          </TabPane>
        </Tabs>
      </Form>
    </Modal>
  );
};

export default Form.create()(PatientWindow);
