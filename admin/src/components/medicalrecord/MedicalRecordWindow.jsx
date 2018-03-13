import React from 'react';
import { Modal, Form, Input, InputNumber, Select, DatePicker, Tabs, Row, Col } from 'antd';
import moment from 'moment';
import PatientSelect from '../patient/PatientSelect';

const FormItem = Form.Item;
const Option = Select.Option;
const TabPane = Tabs.TabPane;

const checkPatient = (rule, value, callback) => {
  if (value.id > 0) {
    callback();
    return;
  }
  callback('Please input the patient');
};

const MedicalRecordWindow = ({ visible, onCancel, onCreate, form, medicalRecord }) => {
  const { getFieldDecorator } = form;
  const patient = medicalRecord.patient && medicalRecord.patient.id > 0 ?
    medicalRecord.patient : { id: undefined };
  return (
    <Modal
      visible={visible}
      title="MedicalRecord"
      okText="Save"
      onCancel={onCancel}
      onOk={onCreate}
      width={600}
      wrapClassName="vertical-center-modal"
      bodyStyle={{ paddingTop: 0, paddingBottom: 5, height: 525 }}
    >
      <Form layout="vertical">
        <Tabs defaultActiveKey="1">
          <TabPane tab="General" key="1">
            <FormItem label="Code">
              {getFieldDecorator('code', {
                initialValue: medicalRecord.code,
                rules: [
                  { required: true, message: 'Please input the code' },
                ],
              })(
                <Input />,
              )}
            </FormItem>
            <FormItem label="Patient">
              {getFieldDecorator('patient', {
                initialValue: patient,
                rules: [
                  { required: true, validator: checkPatient },
                ],
              })(
                <PatientSelect />,
              )}
            </FormItem>
            <FormItem label="Record Date">
              {getFieldDecorator('recordDate', {
                initialValue: moment(medicalRecord.recordDate),
                rules: [
                  { required: true, message: 'Please input record date' },
                ],
              })(
                <DatePicker />,
              )}
            </FormItem>
          </TabPane>
          <TabPane tab="OD" key="2">
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Corneal Diameter">
                  {getFieldDecorator('cornealDiameter', {
                    initialValue: medicalRecord.cornealDiameter,
                  })(
                    <Select>
                      <Option value="H">Horizontal</Option>
                      <Option value="V">Vertical</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row>
              <Col span={12}>
                <FormItem label="Intraocular Pressure">
                  {getFieldDecorator('intraocularPressure', {
                    initialValue: medicalRecord.intraocularPressure,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="Axial Length">
                  {getFieldDecorator('axialLength', {
                    initialValue: medicalRecord.axialLength,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Refraksi">
                  {getFieldDecorator('refraksi', {
                    initialValue: medicalRecord.refraksi,
                  })(
                    <Select>
                      <Option value="S">Spheris</Option>
                      <Option value="C">Cylindar</Option>
                      <Option value="A">Axis</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="Axis">
                  {getFieldDecorator('axis', {
                    initialValue: medicalRecord.axis,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="IOL Type">
                  {getFieldDecorator('iOLType', {
                    initialValue: medicalRecord.iOLType,
                  })(
                    <Select>
                      <Option value="F">Foldable</Option>
                      <Option value="P">PMMA</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="IOL Power">
                  {getFieldDecorator('iOLPower', {
                    initialValue: medicalRecord.iOLPower,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Keratometri">
                  {getFieldDecorator('keratometri', {
                    initialValue: medicalRecord.keratometri,
                  })(
                    <Select>
                      <Option value="K1">K1</Option>
                      <Option value="K2">K2</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Pre Op Visual Acuity">
                  {getFieldDecorator('preOpVisualAcuity', {
                    initialValue: medicalRecord.preOpVisualAcuity,
                  })(
                    <Select>
                      <Option value="OD">OD</Option>
                      <Option value="OS">OS</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="Post Op Visual Acuity">
                  {getFieldDecorator('postOpVisualAcuity', {
                    initialValue: medicalRecord.postOpVisualAcuity,
                  })(
                    <Select>
                      <Option value="OD">OD</Option>
                      <Option value="OS">OS</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
            </Row>
          </TabPane>

          <TabPane tab="OS" key="3">
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Corneal Diameter">
                  {getFieldDecorator('cornealDiameter2', {
                    initialValue: medicalRecord.cornealDiameter2,
                  })(
                    <Select>
                      <Option value="H">Horizontal</Option>
                      <Option value="V">Vertical</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row>
              <Col span={12}>
                <FormItem label="Intraocular Pressure">
                  {getFieldDecorator('intraocularPressure2', {
                    initialValue: medicalRecord.intraocularPressure2,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="Axial Length">
                  {getFieldDecorator('axialLength2', {
                    initialValue: medicalRecord.axialLength2,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Refraksi">
                  {getFieldDecorator('refraksi2', {
                    initialValue: medicalRecord.refraksi2,
                  })(
                    <Select>
                      <Option value="S">Spheris</Option>
                      <Option value="C">Cylindar</Option>
                      <Option value="A">Axis</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="Axis">
                  {getFieldDecorator('axis2', {
                    initialValue: medicalRecord.axis2,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="IOL Type">
                  {getFieldDecorator('iOLType2', {
                    initialValue: medicalRecord.iOLType2,
                  })(
                    <Select>
                      <Option value="F">Foldable</Option>
                      <Option value="P">PMMA</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="IOL Power">
                  {getFieldDecorator('iOLPower2', {
                    initialValue: medicalRecord.iOLPower2,
                  })(
                    <InputNumber step={0.1} style={{ width: '60%' }} />,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Keratometri">
                  {getFieldDecorator('keratometri2', {
                    initialValue: medicalRecord.keratometri2,
                  })(
                    <Select>
                      <Option value="K1">K1</Option>
                      <Option value="K2">K2</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
            </Row>
            <Row gutter={15}>
              <Col span={12}>
                <FormItem label="Pre Op Visual Acuity">
                  {getFieldDecorator('preOpVisualAcuity2', {
                    initialValue: medicalRecord.preOpVisualAcuity2,
                  })(
                    <Select>
                      <Option value="OD">OD</Option>
                      <Option value="OS">OS</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
              <Col span={12}>
                <FormItem label="Post Op Visual Acuity">
                  {getFieldDecorator('postOpVisualAcuity2', {
                    initialValue: medicalRecord.postOpVisualAcuity2,
                  })(
                    <Select>
                      <Option value="OD">OD</Option>
                      <Option value="OS">OS</Option>
                    </Select>,
                  )}
                </FormItem>
              </Col>
            </Row>
          </TabPane>
        </Tabs>
      </Form>
    </Modal>
  );
};

export default Form.create()(MedicalRecordWindow);
