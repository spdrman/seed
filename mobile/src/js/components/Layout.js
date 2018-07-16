import React from "react";
import ReactDOM from "react-dom";

export default class Layout extends React.Component {
    constructor() {
        super();
        this.state = {
            name: "Will"
        };
    }
    render() {
        return (
            <h1>It's {this.state.name}</h1>
        );
    }
}