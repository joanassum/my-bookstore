import React from 'react'
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import AddShoppingCartIcon from '@material-ui/icons/AddShoppingCart';
import Typography from '@material-ui/core/Typography';

class BookItem extends React.Component {
    constructor(props) {
        super(props)

        this.state = { 
            id: this.props.id, 
            author: this.props.author, 
            title: this.props.title,
            price: this.props.price
        }
    }

    render() {
        return (
            <Card variant="outlined" style={{maxWidth: 200, textAlign: "center"}}>
                <CardContent>
                    <img src={`http://localhost:8080/book/${this.state.id}/image`} style={{width: "110px", height:"160px",objectFit: "cover"}} alt={this.state.title} />
                    <Typography gutterBottom>
                        {this.state.title}
                    </Typography>
                    <Typography color="textSecondary" variant="body2" component="p">
                        {this.state.author}
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button size="small">More</Button>
                    <Typography color="textSecondary" variant="body2" component="p">
                        £{this.state.price}
                    </Typography>
                    <Button variant="contained" color="primary" size="small"><AddShoppingCartIcon/></Button>
                </CardActions>
            </Card>
        )
    }
}

export default BookItem;