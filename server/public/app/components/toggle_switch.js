import React, { Component } from 'react';
import Toggle from 'material-ui/Toggle';
import axios from 'axios';
const styles = {
  block: {
    maxWidth: 250,
  },
  toggle: {
    marginBottom: 16,
    margin: '0 auto',
    width:'100%',
    display: 'inline-block'
  },
};

export default class ToggleSwitch extends Component{
    constructor(props){
      super(props);
      this.state = {'Toggled':this.props.checked};
      this.onToggled = this.onToggled.bind(this);
    }
    onToggled(event, checked){
      console.log(checked);
      if (checked){
        axios.get(`http://home.suyash.io/devices/${this.props.endpoint}/${this.props.func}/on`);
      }else{
        axios.get(`http://home.suyash.io/devices/${this.props.endpoint}/${this.props.func}/off`);
      }
      this.setState({'Toggled':checked});
      console.log('Toggle State is',checked);
      this.props.iChanged(checked);
    }
    render(){
      //this.setState({'Toggled':this.props.checked});
      console.log(this.props.checked);
      return (
        <div className="row">
        <div className="text-center col-sm-4 col-sm-offset-4" style={{display:'flex',"justifyContent":'center', "alignItems":"center"}}>
        <div>
        <Toggle
          styles={styles.toggle}
          onToggle={this.onToggled}
          toggled = {this.props.checked}
        />
        </div>
        </div>
        </div>

      )
    }
}
