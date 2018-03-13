import React from 'react';
import { Select } from 'antd';

const Option = Select.Option;

const RoleSelect = ({ roles, value, onSelect }) => (
  <Select
    placeholder="Select Role"
    style={{ width: 120 }}
    onSelect={onSelect}
    value={value}
  >
    {roles.map(role => (
      <Option value={role.id}>{role.name}</Option>
    ))}
  </Select>
);

export default RoleSelect;
