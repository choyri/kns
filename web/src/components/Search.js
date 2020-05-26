import React from 'react'
import {Input, Button} from 'element-react'

class Search extends React.Component {
    keyword = ''

    handleKeywordChange = keyword => {
        this.keyword = keyword
        this.props.onKeywordChange(keyword)
    }

    handleButtonClick = () => {
        this.props.onSearch(this.keyword)
    }

    render() {
        const {className} = this.props

        return (
            <Input
                className={`container keyword-input ${className}`}
                placeholder="输入关键词，用空格分割"
                size="large"
                autoFocus
                trim
                onChange={this.handleKeywordChange}
                append={
                    <Button icon="search" onClick={this.handleButtonClick}>搜索</Button>
                }
            />
        )
    }
}

export default Search
