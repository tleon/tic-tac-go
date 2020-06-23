import React from 'react';
import './App.css';
import Axios from "axios";

class App extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            grid: null
        }
    }

    componentDidMount() {
        Axios
            .get("http://localhost:8080/get-grid")
            .then(r => {
                this.setState({
                    grid: r.data
                })
            })
    }

    render() {
        const { grid } = this.state
        console.log("grid ", grid)
        return (
            <div>
                {grid != null
                    ? (
                        grid.map((idx, line) => {
                            console.log(idx, line)
                        })
                    )
                    : "there is no grid"
                }
            </div>
        )
    }
}

export default App;
