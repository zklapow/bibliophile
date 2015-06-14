
export default class BookItem extends React.Component {
    render() {
        return (
            <div>
                <div className="left">
                    <img className="cover-image" src="http://lukeloaghan.files.wordpress.com/2011/11/5281243873_c8c83b865a.jpg" width="100"/>
                </div>
                <div>
                    <h3>{this.props.book.Title}</h3>
                    <h4>{this.props.book.Author}</h4>
                </div>
            </div>
        )
    }
}