import React from 'react';
import Axios from "axios";
import Cell from "./Cell";
import {Col, Container, Row} from "react-bootstrap";

class App extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            grid: null
        }
    }

    handleClick(row, col) {
        Axios
            .post(
                "http://localhost:8080/clicked",
                {
                    row,
                    col
                }
            ).then(r => {
                this.setState({
                    grid: r.data
                })
            })
        ;
    }

    componentDidMount() {
        Axios
            .get("http://localhost:8080/get-grid")
            .then(r => {
                this.setState({
                    grid: r.data
                })
            })
        ;
    }

    render() {
        const { grid } = this.state

        return (
            <div className={grid}>
                {grid != null
                    ? (
                        <div>
                            <br/>
                            <br/>
                            <Container>
                                <Row>
                                    <Col>
                                        <div onClick={() => this.handleClick(0, 0)}>
                                            <Cell isChecked={grid[0][0].IsChecked} byPlayer={grid[0][0].ByPlayer} />
                                        </div>
                                    </Col>
                                    <Col>
                                        <div onClick={() => this.handleClick(0, 1)}>
                                            <Cell isChecked={grid[0][1].IsChecked} byPlayer={grid[0][1].ByPlayer} />
                                        </div>
                                    </Col>
                                    <Col>
                                        <div onClick={() => this.handleClick(0, 2)}>
                                            <Cell isChecked={grid[0][2].IsChecked} byPlayer={grid[0][2].ByPlayer} />
                                        </div>
                                    </Col>
                                </Row>
                                <Row>
                                    <Col>
                                        <div onClick={() => this.handleClick(1, 0)}>
                                            <Cell isChecked={grid[1][0].IsChecked} byPlayer={grid[1][0].ByPlayer} />
                                        </div>
                                    </Col>
                                    <Col>
                                        <div onClick={() => this.handleClick(1, 1)}>
                                            <Cell isChecked={grid[1][1].IsChecked} byPlayer={grid[1][1].ByPlayer} />
                                        </div>
                                    </Col>
                                    <Col>
                                        <div onClick={() => this.handleClick(1, 2)}>
                                            <Cell isChecked={grid[1][2].IsChecked} byPlayer={grid[1][2].ByPlayer} />
                                        </div>
                                    </Col>
                                </Row>
                                <Row>
                                    <Col>
                                        <div onClick={() => this.handleClick(2, 0)}>
                                            <Cell isChecked={grid[2][0].IsChecked} byPlayer={grid[2][0].ByPlayer} />
                                        </div>
                                    </Col>
                                    <Col>
                                        <div onClick={() => this.handleClick(2, 1)}>
                                            <Cell isChecked={grid[2][1].IsChecked} byPlayer={grid[2][1].ByPlayer} />
                                        </div>
                                    </Col>
                                    <Col>
                                        <div onClick={() => this.handleClick(2, 2)}>
                                            <Cell isChecked={grid[2][2].IsChecked} byPlayer={grid[2][2].ByPlayer} />
                                        </div>
                                    </Col>
                                </Row>
                            </Container>
                        </div>
                    )
                    : "there is no grid"
                }
            </div>
        )
    }
}

export default App;
