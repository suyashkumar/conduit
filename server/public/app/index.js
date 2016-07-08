import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import ToggleSwitch from './components/toggle_switch';
import ToggleButtons from './components/toggle_buttons';
class App extends Component{
	constructor(props){
		super(props);
		this.state={'ledOn':false}
		this.stateChanged = this.stateChanged.bind(this);
	}
	stateChanged(checked){
		this.setState({'ledOn':checked})
		console.log('state changed');
	}
	render(){
		return (
			<div>
				<div className="well row text-center" style={{"margin-top":"20px"}}>
					<h1>Home Auto Test</h1>
				</div>

				<ToggleButtons
					endpoint="suyash"
					func="led"
					iChanged={this.stateChanged}/>


				<MuiThemeProvider>
					<ToggleSwitch
						endpoint="suyash"
						func="led"
						checked={this.state.ledOn}
						iChanged={this.stateChanged}/>

				</MuiThemeProvider>



	</div>
		)
	}
}

ReactDOM.render(<App />, document.getElementById('app'));
