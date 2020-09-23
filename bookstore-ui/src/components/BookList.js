import React from 'react'
import GridList from '@material-ui/core/GridList';
import BookItem from './BookItem'

class BookList extends React.Component {
    constructor(props) {
        super(props)
        this.state = { books: [] }

        this.getBooksList = this.getBooksList.bind(this)
    }

    componentDidMount() {
        this.getBooksList()
    }

    async getBooksList() {
        try {
            await fetch("http://localhost:8080/books")
            .then(res => res.json())
            .then(data => {
                console.log(data)
                this.setState({books: data})
            });
        } catch (e) {
            console.error(e);
        }
    }

    render() {
        return (
            <div>
                <br></br>
                <GridList cols={5} cellHeight={'auto'}>
                    {this.state.books.map((book) => (
                        <BookItem 
                            key={book.id}
                            id={book.id}
                            author={book.author}
                            title={book.title}
                            price={book.price}
                        />
                    ))}
                </GridList>
            </div>
        )
    }

    
}

export default BookList;