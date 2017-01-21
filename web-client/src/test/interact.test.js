import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import ReactTestUtils from 'react-addons-test-utils'
import axios from '../axios-auth';
import Interact from '../components/interact'; 
import sinon from 'sinon';

it('Renders without crashing', () => { 
  const div = document.createElement('div');
  ReactDOM.render(<Interact />, div);
});


it('Axios Call Made on button click', () => { 
	const c = ReactTestUtils.renderIntoDocument(<Interact />);
	const mockFunc = jest.fn(); 
	const p = new Promise((resolve, reject) => resolve());
	const spy = sinon.stub(axios, 'get').returns(p);
	const button = ReactTestUtils.findRenderedDOMComponentWithClass(c, 'button') 
	ReactTestUtils.Simulate.click(button); 
	expect(spy.called).toBe(true);
	spy.restore(); 
});


