import React, { Component } from 'react';
import Toggle from 'material-ui/Toggle';
import axios from 'axios';
const styles = {
  block: {
    maxWidth: 250,
  },
  toggle: {
    marginBottom: 16,
    maxWidth:300,
    minWidth:200
  },
};

export default class ToggleSwitch extends Component{
    constructor(props){
      super(props);
      this.onToggled = this.onToggled.bind(this);
    }
    onToggled(event, checked){
      console.log(checked);
      if (checked){
        axios.get(`http://home.suyash.io/${this.props.endpoint}/${this.props.func}/on`);
      }else{
        axios.get(`http://home.suyash.io/${this.props.endpoint}/${this.props.func}/off`);
      }

    }
    render(){
      return (
        <div className="text-center col-sm-4 ">
        <Toggle
        styles={styles.toggle}
        onToggle={this.onToggled}
        />
        </div>

      )
    }
}
