import React from "react";
import {Card, Divider, Image, Table} from "semantic-ui-react";
import eichel from "../public/eichel.jpg";
import sabres from "../public/sabres-icon.png";


// noinspection SpellCheckingInspection
export default class Player extends React.Component {
    constructor(props) {
        super(props);
        ({
            id: this.id,
            full_name: this.name,
            height: this.height,
            weight: this.weight,
            primary_position: this.pos,
            jersey_number: this.num,
            statistics: this.statistics,
        } = props);
    }

    calcHeight() {
        let ft = Math.floor(this.height/12)
        let inch = Math.floor(this.height%12)
        return (ft.toString() + "'" + inch.toString() + "\"")
    }

    stats() {
        let stats = this.statistics["total"];
        return <Table padded>
            <Table.Header>
                <Table.Row textAlign='center'>
                    <Table.HeaderCell><abbr title={"Games Played"}>GP</abbr></Table.HeaderCell>
                    <Table.HeaderCell>Goals</Table.HeaderCell>
                    <Table.HeaderCell>Assists</Table.HeaderCell>
                    <Table.HeaderCell>Points</Table.HeaderCell>
                </Table.Row>
            </Table.Header>
            <Table.Body>
                <Table.Row textAlign='center'>
                    <Table.Cell>{stats.games_played}</Table.Cell>
                    <Table.Cell>{stats.goals}</Table.Cell>
                    <Table.Cell>{stats.assists}</Table.Cell>
                    <Table.Cell>{stats.points}</Table.Cell>
                </Table.Row>
            </Table.Body>
        </Table>
    }

    render() {
        return <Card key={this.id} fluid>
            <Card.Content>
                <Card.Header>
                    <Image size="tiny" src={eichel}/>
                    <Divider hidden/>
                    {this.name} | #{this.num}
                </Card.Header>
                <Card.Meta textAlign='center'>
                    {this.pos} | {this.calcHeight()} | {this.weight} lb
                </Card.Meta>
                <br/>
                <Card.Content extra>
                    <Image src={sabres} size='mini' rounded/>
                    {this.stats()}
                </Card.Content>
            </Card.Content>
        </Card>
    }

}