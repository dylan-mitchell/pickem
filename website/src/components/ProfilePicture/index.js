import React from "react";
import Image from "react-bulma-components/lib/components/image";

class ProfilePicture extends React.Component {
  getPicture = () => {
    if (!this.props.imageURL) {
      return (
        "https://api.adorable.io/avatars/285/" + this.props.username + ".png"
      );
    } else {
      return this.props.imageURL;
    }
  };

  render() {
    return <Image src={this.getPicture()} size={this.props.size} rounded />;
  }
}

export default ProfilePicture;
