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
  link: "",
  error: null
};
class UpdateProfilePictureFormBase extends Component {
  constructor(props) {
    super(props);
    this.state = { ...INITIAL_STATE };
  }
  onSubmit = event => {
    const { link } = this.state;
    this.props.firebase
      .doUpdateProfilePicture(link)
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
    const { link, error } = this.state;
    const lastPart = link.split(".").pop();
    const isValidLink =
      lastPart === "jpg" ||
      lastPart === "jpeg" ||
      lastPart === "png" ||
      lastPart === "" ||
      lastPart === "gif";
    return (
      <Box>
        <Level renderAs="nav">
          <Level.Item>
            <Field>
              <Control>
                <Input
                  onChange={this.onChange}
                  name="link"
                  type="link"
                  placeholder="Link to Image"
                  value={link}
                  fullwidth
                />
              </Control>
            </Field>
          </Level.Item>
          <Level.Item>
            <Button
              disabled={!isValidLink}
              type="submit"
              onClick={this.onSubmit}
              color="success"
              size="medium"
              fullwidth
            >
              Update Profile Picture
            </Button>
          </Level.Item>
        </Level>
        <p>Please use URL's that link to JPG, PNG, or GIF resources</p>
        {error && <p style={helpStyle}>{error.message}</p>}
      </Box>
    );
  }
}

const helpStyle = {
  color: "red"
};

const UpdateProfilePictureForm = compose(
  withRouter,
  withFirebase
)(UpdateProfilePictureFormBase);

export default withFirebase(UpdateProfilePictureForm);
