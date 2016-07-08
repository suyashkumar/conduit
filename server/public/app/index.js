import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import ToggleSwitch from './components/toggle_switch';

class App extends Component{
	constructor(props){
		super(props);

	}
	render(){
		return (
			<div>
			<MuiThemeProvider>
				<ToggleSwitch endpoint="suyash" func="led"/>
			</MuiThemeProvider>
		</div>

		)
	}
}

ReactDOM.render(<App />, document.getElementById('app'));
