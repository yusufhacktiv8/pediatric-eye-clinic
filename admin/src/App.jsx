import React from 'react';

import {
  HashRouter as Router,
  Route,
} from 'react-router-dom';
import 'antd/dist/antd.css';

import './App.css';

import Workspace from './components/workspace/Workspace';
import DashboardPage from './components/dashboard/DashboardPage';
import UserPage from './components/user/UserPage';
import RolePage from './components/role/RolePage';
import InsurancePage from './components/insurance/InsurancePage';
import OccupationPage from './components/occupation/OccupationPage';
import DiseasePage from './components/disease/DiseasePage';
import PatientPage from './components/patient/PatientPage';
import MedicalRecordPage from './components/medicalrecord/MedicalRecordPage';
import LoginForm from './components/login/LoginForm';

export default () => {
  const token = window.sessionStorage.getItem('token');

  let componentToRender = (
    <Router>
      <Workspace>
        <Route exact path="/" component={DashboardPage} />
        <Route exact path="/users" component={UserPage} />
        <Route exact path="/roles" component={RolePage} />
        <Route exact path="/insurances" component={InsurancePage} />
        <Route exact path="/occupations" component={OccupationPage} />
        <Route exact path="/diseases" component={DiseasePage} />
        <Route exact path="/patients" component={PatientPage} />
        <Route exact path="/medicalrecords" component={MedicalRecordPage} />
      </Workspace>
    </Router>
  );

  // if (!token) {
  //   componentToRender = (
  //     <LoginForm />
  //   );
  // }

  return (
    <div className="App">
      {componentToRender}
    </div>
  );
};
