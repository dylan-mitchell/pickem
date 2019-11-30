import React, { Component } from "react";
import { withFirebase } from "../Firebase";
import {
  Field,
  Control,
  Input
} from "react-bulma-components/lib/components/form";
import Button from "react-bulma-components/lib/components/button";
import Box from "react-bulma-components/lib/components/box";
import Level from "react-bulma-components/lib/components/level";
import * as ROUTES from "../../constants/routes";
import { withRouter } from "react-router-dom";
import { compose } from "recompose";

const INITIAL_STATE = {
  username: "",
  error: null
};
class UpdateUsernameFormBase extends Component {
  constructor(props) {
    super(props);
    this.state = { ...INITIAL_STATE };
  }
  onSubmit = event => {
    const { username } = this.state;
    this.props.firebase
      .doUpdateUsername(username)
      .then(() => {
        this.setState({ ...INITIAL_STATE });
        this.props.history.push(ROUTES.ACCOUNT);
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
    const { username, error } = this.state;

    return (
      <Box>
        <Level renderAs="nav">
          <Level.Item>
            <Field>
              <Control>
                <Input
                  onChange={this.onChange}
                  name="username"
                  placeholder="Username"
                  value={username}
                  fullwidth
                />
              </Control>
            </Field>
          </Level.Item>
          <Level.Item>
            <Button
              disabled={username === ""}
              type="submit"
              onClick={this.onSubmit}
              color="success"
              size="medium"
              fullwidth
            >
              Update Username
            </Button>
          </Level.Item>
        </Level>
        {error && <p style={helpStyle}>{error.message}</p>}
      </Box>
    );
  }
}

const helpStyle = {
  color: "red"
};

const UpdateUsernameForm = compose(
  withRouter,
  withFirebase
)(UpdateUsernameFormBase);

export default withFirebase(UpdateUsernameForm);
