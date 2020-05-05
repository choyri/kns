import React from 'react'
import {Input, Button} from 'element-react'

class Search extends React.Component {
    state = {
        keyword: '',
    }

    handleKeywordChange = keyword => {
        this.setState({keyword})
    }

    render() {
        const {keyword} = this.state
        const {className, onSearch} = this.props

        return (
            <Input
                className={`container keyword-input ${className}`}
                placeholder="输入关键词，用空格分割"
                size="large"
                autoFocus
                trim
                value={keyword}
                onChange={this.handleKeywordChange}
                append={<Button icon="search" onClick={onSearch.bind(this, keyword)}>搜索</Button>}
            />
        )
    }
}

export default Search
