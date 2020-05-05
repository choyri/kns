import React from 'react'
import {MessageBox} from 'element-react'
import Header from './components/Header'
import Query from './components/Query'
import Import from './components/Import'
import {visitBackend} from './fetch'

import './App.scss'

class App extends React.Component {
    state = {
        activeMenu: '',
    }

    menus = [
        {index: 'query', name: '查询'},
        {index: 'import', name: '导入'},
    ]

    constructor(props) {
        super(props)
        this.state.activeMenu = this.menus[0].index
    }

    componentDidMount() {
        visitBackend().catch(err => {
            MessageBox.alert(`请检查服务端是否正常开启（${err}）`, '服务端未正常工作', {type: 'error'})
        })
    }

    handleMenuSelect = activeMenu => {
        this.setState({activeMenu})
    }

    render() {
        const {activeMenu} = this.state

        return (
            <>
                <Header
                    menus={this.menus}
                    defaultActiveMenu={activeMenu}
                    handleMenuSelect={this.handleMenuSelect}
                />
                {activeMenu === 'query' && <Query/>}
                {activeMenu === 'import' && <Import/>}
            </>
        )
    }
}

export default App
