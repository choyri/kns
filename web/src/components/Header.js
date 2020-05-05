import React from 'react'
import {Menu} from 'element-react'

class Header extends React.Component {
    render() {
        const {defaultActiveMenu, handleMenuSelect} = this.props

        return (
            <header>
                <div className='container'>
                    <Menu
                        mode='horizontal'
                        defaultActive={defaultActiveMenu}
                        onSelect={handleMenuSelect.bind(this)}
                    >
                        {this.props.menus.map(element =>
                            <Menu.Item key={element.index} index={element.index}>{element.name}</Menu.Item>
                        )}
                    </Menu>
                </div>
            </header>
        )
    }
}

export default Header
