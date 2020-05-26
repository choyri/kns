import React from 'react'
import {Button, Pagination} from 'element-react'
import Search from './Search'
import Order from './Order'
import {search, exportData} from '../fetch'

import './Query.scss'

class Query extends React.Component {
    state = {
        isSearched: false,
        searchResult: {
            lists: [],
            total: 0,
            page: 1,
            perPage: 0,
        },
        exportData: [],
    }

    keyword = ''

    disposeSearchResultClassName = () => {
        let ret = ['search-result']

        let exportDataLen = this.state.exportData.length
        if (exportDataLen > 3) {
            exportDataLen = 3
        }

        if (exportDataLen > 0) {
            ret.push(`for-export-${exportDataLen}`)
        }

        return ret.join(' ')
    }

    handleKeywordChange = keyword => {
        this.keyword = keyword
    }

    handleSearch = keyword => {
        keyword || (keyword = this.keyword)

        if (keyword === '') {
            let searchResult = this.state.searchResult
            searchResult.lists = []
            searchResult.total = 0
            this.setState({searchResult})
            return
        }

        const {page, perPage} = this.state.searchResult

        search(keyword, page, perPage).then(({results, headers}) => {
            const searchResult = {
                lists: results,
                total: parseInt(headers.get('X-Total')),
                page: parseInt(headers.get('X-Page')),
                perPage: parseInt(headers.get('X-Per-Page')),
            }

            let data = {searchResult}

            if (!this.state.isSearched) {
                data.isSearched = true
            }

            this.setState(data)
        })
    }

    handlePerPageChange = perPage => {
        let searchResult = this.state.searchResult
        searchResult.perPage = perPage
        this.setState({searchResult}, () => {
            this.handleSearch()
        })
    }

    handleCurrentPageChange = page => {
        let searchResult = this.state.searchResult
        searchResult.page = page
        this.setState({searchResult}, () => {
            this.handleSearch()
        })
    }

    handlePitch = row => {
        this.setState({
            exportData: [...this.state.exportData, row],
        })
    }

    handleRemove = index => {
        let {exportData} = this.state
        exportData.splice(index, 1)

        this.setState({exportData: [...exportData]})
    }

    handleExport = () => {
        let ids = []

        for (const element of this.state.exportData) {
            ids.push(element.id)
        }

        exportData(ids)
    }

    render() {
        const {isSearched, searchResult, exportData} = this.state

        return (
            <main>
                <Search
                    className={isSearched ? 'sideways' : ''}
                    onKeywordChange={this.handleKeywordChange}
                    onSearch={this.handleSearch}
                />

                {isSearched && <div className={this.disposeSearchResultClassName()}>
                    <Order
                        data={searchResult.lists}
                        onPitch={this.handlePitch}
                        canPitch
                    />
                    <Pagination
                        layout="sizes, prev, pager, next, total"
                        total={searchResult.total}
                        pageSize={searchResult.perPage}
                        currentPage={searchResult.page}
                        onSizeChange={this.handlePerPageChange}
                        onCurrentChange={this.handleCurrentPageChange}
                    />
                </div>}

                <div className={`export-data ${exportData.length > 0 ? 'show' : ''}`}>
                    <div className="header">
                        <span>目标数据</span>
                        <Button
                            type="primary" size="small" icon="upload2"
                            onClick={this.handleExport}
                        >导出</Button>
                    </div>
                    <Order
                        maxHeight={160}
                        data={exportData}
                        onRemove={this.handleRemove}
                        canRemove
                    />
                </div>
            </main>
        )
    }
}

export default Query
