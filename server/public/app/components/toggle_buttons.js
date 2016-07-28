import React, { Component } from 'react';
import axios from 'axios';
export default class ToggleSwitch extends Component{
    constructor(props){
      super(props);
    }
    onToggle(checked){
      if(checked){
        axios.get(`http://home.suyash.io/devices/${this.props.endpoint}/${this.props.func}/on`);
      }
      else{
        axios.get(`http://home.suyash.io/devices/${this.props.endpoint}/${this.props.func}/off`);

      }
      this.props.iChanged(checked);
    }
    render(){
      return (
        <div className="text-center row">
          <div className="col-sm-4"><h2>{this.props.endpoint} | {this.props.func}</h2></div>
          <div className='col-sm-8'>
            <button style={{"margin-right":"10px"}} className="btn btn-lg btn-primary" onClick={()=>{this.onToggle(true)}}>On</button>
            <button className="btn btn-lg btn-danger" onClick={()=>{this.onToggle(false)}}>Off</button>
          </div>
        </div>
      )
    }
}
