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
			<div >


				<ToggleButtons
					endpoint="suyash"
					func="led" />
		</div>

		)
	}
}

ReactDOM.render(<App />, document.getElementById('app'));
