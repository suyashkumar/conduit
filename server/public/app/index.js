import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import ToggleSwitch from './components/toggle_switch';
import ToggleButtons from './components/toggle_buttons';
class App extends Component{
	constructor(props){
		super(props);

	}
	render(){
		return (
			<div>
				<div className="well text-center" style={{"margin-top":"20px"}}>
					<h1>Home Auto Test</h1>
				</div>
			<div className="row">

				<ToggleButtons
					endpoint="suyash"
					func="led" />
			</div>

	</div>
		)
	}
}

ReactDOM.render(<App />, document.getElementById('app'));
