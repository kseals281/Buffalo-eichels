import React from 'react';
import sabres from "./sabres-icon.png";
import './App.css';
import Player from './Player'
import Goalie from './Goalie'
import Axios from 'axios';
import {Card, Container, Divider, HeaderContent} from "semantic-ui-react";

const endpoint = "http://localhost:8080";

class BuffaloEichels extends React.Component {
    constructor(props) {
        super(props);
        this.state = {forwards: [], defencemen: [], goalies: []}
    }

    forwards() {
        Axios.get(endpoint + "/api/forwards").then(res => {
            console.log("Getting forwards")
            console.log(res.data)
            this.setState({
                    forwards: res.data.map(item => {
                        return <Player {...item} />
                    })
                }
            )
        })
    }

    defence() {
        Axios.get(endpoint + "/api/defencemen").then(res => {
            console.log("Getting defensemen")
            console.log(res.data)
            this.setState({
                    defencemen: res.data.map(item => {
                        return <Player {...item} />
                    })
                }
            )
        })
    }

    goalies() {
        Axios.get(endpoint + "/api/goalies").then(res => {
            console.log("Getting goalies")
            console.log(res.data)
            this.setState({
                    goalies: res.data.map(item => {
                        console.log(item)
                        return <Goalie {...item} />
                    })
                }
            )
        })
    }

    componentDidMount() {
        console.log("Getting roster");
        this.forwards()
        this.defence()
        this.goalies()
    }

    render() {
        return (
            <div className="App">
                <header className="Eichel-header">
                    <HeaderContent className='Title-text-left'>BUFFALO</HeaderContent>
                    <img src={sabres} className="App-logo" alt="logo" sizes={"medium"}/>
                    <HeaderContent className="Title-text-right">EICHELS</HeaderContent>
                </header>
                <Divider/>
                <Container>
                    <Card.Group itemsPerRow={3}>
                        {this.state.forwards}
                    </Card.Group>
                    <Card.Group centered itemsPerRow={2}>
                        {this.state.defencemen}
                    </Card.Group>
                    {/*<Card.Group centered itemsPerRow={1}>*/}
                    {/*    {this.state.goalies}*/}
                    {/*</Card.Group>*/}
                </Container>
                <Divider/>
                <footer className="App-footer">
                    I own nothing
                </footer>
            </div>
        );
    }
}

export default BuffaloEichels;
