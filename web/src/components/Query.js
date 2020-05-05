import React from 'react'
import {Button} from 'element-react'
import Search from './Search'
import Order from './Order'
import {search, exportData} from '../fetch'

import './Query.scss'

class Query extends React.Component {
    state = {
        isSearched: false,
        searchResults: [],
        exportData: [],
    }

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

    handleSearch = keyword => {
        if (keyword === '') {
            this.setState({searchResults: []})
            return
        }

        search(keyword).then(results => {
            this.setState({searchResults: results})
            if (!this.state.isSearched) {
                this.setState({isSearched: true})
            }
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
        const {isSearched, searchResults, exportData} = this.state

        return (
            <main>
                <Search
                    className={isSearched ? 'sideways' : ''}
                    onSearch={this.handleSearch}
                />

                {isSearched && <div className={this.disposeSearchResultClassName()}>
                    <Order
                        data={searchResults}
                        onPitch={this.handlePitch}
                        canPitch
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
