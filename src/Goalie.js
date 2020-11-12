import React from "react";
import {Card, Divider, Image, Table} from "semantic-ui-react";
import eichel from "./eichel.jpg";
import sabres from "./sabres-icon.png";


export default class Goalie extends React.Component {
    constructor(props) {
        super(props);
        ({
            name: this.name,
            height: this.height,
            weight: this.weight,
            pos: this.pos,
            num: this.num,
            games: this.games,
            started: this.started,
            wins: this.wins,
            losses: this.losses,
            ot_losses: this.ot,
            save_pct: this.savePct,
            goals_against: this.gaa,
        } = props);
    }

    goalieStats() {
        console.log(this.savePct)
        return <Table padded>
            <Table.Header>
                <Table.Row>
                    <Table.HeaderCell><abbr title='Games Played'>GP</abbr></Table.HeaderCell>
                    <Table.HeaderCell><abbr title='Started'>GS</abbr></Table.HeaderCell>
                    <Table.HeaderCell><abbr title='Wins'>W</abbr></Table.HeaderCell>
                    <Table.HeaderCell><abbr title='Losses'>L</abbr></Table.HeaderCell>
                    <Table.HeaderCell><abbr title='Overtime Losses'>OTL</abbr></Table.HeaderCell>
                    <Table.HeaderCell><abbr title='Save Percentage'>sv %</abbr></Table.HeaderCell>
                    <Table.HeaderCell><abbr title='Goals against average'>GAA</abbr></Table.HeaderCell>
                </Table.Row>
            </Table.Header>
            <Table.Body>
                <Table.Row>
                    <Table.Cell>{this.games}</Table.Cell>
                    <Table.Cell>{this.started}</Table.Cell>
                    <Table.Cell>{this.wins}</Table.Cell>
                    <Table.Cell>{this.losses}</Table.Cell>
                    <Table.Cell>{this.ot}</Table.Cell>
                    <Table.Cell>{this.savePct}</Table.Cell>
                    <Table.Cell>{this.gaa}</Table.Cell>
                </Table.Row>
            </Table.Body>
        </Table>
    }

    render() {
        return <Card fluid>
            <Card.Content>
                <Card.Header>
                    <Image size="tiny" src={eichel}/>
                    <Divider hidden/>
                    Jack Eichel | #{this.num}
                </Card.Header>
                <Card.Meta textAlign='center'>
                    {this.pos} | {this.height} | {this.weight} lb
                </Card.Meta>
                <br/>
                <Card.Content extra>
                    <Image src={sabres} size='mini' rounded/>
                    {this.goalieStats()}
                </Card.Content>
            </Card.Content>
        </Card>
    }

}