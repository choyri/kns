import React from 'react'
import {Button, Table} from 'element-react'

class Order extends React.Component {
    commonColumns = [
        {label: '客户单号', prop: 'customer_order_number'},
        {label: '品牌', prop: 'brand'},
        {label: '订单号码', prop: 'order_number'},
        {label: '序号', prop: 'serial_number', width: 80},
        {label: '品名代码', prop: 'product_name_code'},
        {label: '成分', prop: 'ingredient'},
        {label: '规格', prop: 'specification'},
        {label: '颜色', prop: 'color'},
        {label: '客户版号', prop: 'customer_version_number'},
    ]

    pitchColumn = {
        label: '操作',
        width: 70,
        fixed: 'right',
        render: row => {
            return <Button type="text" size="small" onClick={this.props.onPitch.bind(this, row)}>添加</Button>
        }
    }

    removeColumn = {
        label: '操作',
        width: 70,
        fixed: 'right',
        render: (row, column, index) => {
            return <Button type="text" size="small" onClick={this.props.onRemove.bind(this, index)}>移除</Button>
        }
    }

    render() {
        let columns = [...this.commonColumns]

        if (this.props.canPitch) {
            columns.push(this.pitchColumn)
        }

        if (this.props.canRemove) {
            columns.push(this.removeColumn)
        }

        return (
            <Table
                columns={columns}
                data={this.props.data}
                stripe
                {...this.props}
            />
        )
    }
}

export default Order