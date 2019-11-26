import React, { Component } from "react";
import { withRouter } from "react-router-dom";
import { compose } from "recompose";
import { SignUpLink } from "../SignUp";
import { PasswordForgetLink } from "../PasswordForget";
import { withFirebase } from "../Firebase";
import * as ROUTES from "../../constants/routes";

import Section from "react-bulma-components/lib/components/section";
import Heading from "react-bulma-components/lib/components/heading";
import Container from "react-bulma-components/lib/components/container";
import {
  Field,
  Control,
  Input
} from "react-bulma-components/lib/components/form";
import Button from "react-bulma-components/lib/components/button";
import Columns from "react-bulma-components/lib/components/columns";

const SignInPage = () => (
  <Section>
    <Container>
      <Columns centered>
        <Columns.Column size="half">
          <Heading size={1} style={centerStyle}>
            Sign In
          </Heading>
          <SignInForm />
          <PasswordForgetLink />
          <SignUpLink />
        </Columns.Column>
      </Columns>
    </Container>
  </Section>
);
const INITIAL_STATE = {
  email: "",
  password: "",
  error: null
};

class SignInFormBase extends Component {
  constructor(props) {
    super(props);
    this.state = { ...INITIAL_STATE };
  }

  onSubmit = event => {
    const { email, password } = this.state;
    this.props.firebase
      .doSignInWithEmailAndPassword(email, password)
      .then(() => {
        this.setState({ ...INITIAL_STATE });
        this.props.history.push(ROUTES.HOME);
      })
      .catch(error => {
        this.setState({ error });
      });
    event.preventDefault();
  };
  onChange = event => {
    this.setState({ [event.target.name]: event.target.value });
  };

  render() {
    const { email, password, error } = this.state;
    const isInvalid = password === "" || email === "";
    return (
      <div>
        <Field>
          <Heading size={4}>Email</Heading>
          <Control>
            <Input
              onChange={this.onChange}
              name="email"
              type="email"
              placeholder="Email input"
              value={email}
            />
          </Control>
        </Field>
        <Field>
          <Heading size={4}>Password</Heading>
          <Control>
            <Input
              onChange={this.onChange}
              name="password"
              type="password"
              placeholder="Password"
              value={password}
            />
          </Control>
        </Field>
        {error && <p style={helpStyle}>{error.message}</p>}
        <Button
          disabled={isInvalid}
          type="submit"
          onClick={this.onSubmit}
          color="success"
          size="large"
        >
          Sign In
        </Button>
      </div>
    );
  }
}

const helpStyle = {
  color: "red"
};

const centerStyle = {
  textAlign: "center"
};

const SignInForm = compose(
  withRouter,
  withFirebase
)(SignInFormBase);
export default SignInPage;
export { SignInForm };
