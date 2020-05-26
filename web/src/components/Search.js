import React from 'react'
import {Input, Button} from 'element-react'

class Search extends React.Component {
    keyword = ''

    handleKeywordChange = event => {
        this.keyword = event.target.value
        this.props.onKeywordChange(this.keyword)
    }

    handleKeyPress = event => {
        if (event.key === 'Enter') {
            event.target.blur()
            this.props.onSearch(this.keyword)
        }
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
                onBlur={this.handleKeywordChange}
                onKeyPress={this.handleKeyPress}
                append={
                    <Button icon="search" onClick={this.handleButtonClick}>搜索</Button>
                }
            />
        )
    }
}

export default Search
