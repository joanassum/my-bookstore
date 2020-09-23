import React from 'react'
import { BrowserRouter, Route } from 'react-router-dom';
import BookList from './BookList'

const App = () => {
    return (
        <div className="App">
            <BrowserRouter>
                <div>
                    <h1>My Bookstore</h1>
                    <Route
                        path="/" exact
                        render={(props) => <BookList {...props} />}
                    />
                </div>
            </BrowserRouter>
        </div>
    )
}

export default App;