import React from 'react'
import {MessageBox, Upload, Loading} from 'element-react'
import config from '../config'

import './Import.scss'

class Import extends React.Component {
    state = {
        isUploading: false,
    }

    handleBeforeUpload = () => {
        this.setState({isUploading: true})
    }

    handleUploadSuccess = resp => {
        MessageBox.alert(`共 ${resp.amount} 条数据`, '导入成功', {type: 'success'})
    }

    handleUploadError = err => {
        MessageBox.alert(err.message, '导入失败', {type: 'error'})
    }

    handleUploadChange = () => {
        this.setState({isUploading: false})
    }

    render() {
        return (
            <main>
                <Loading loading={this.state.isUploading}>
                    <Upload
                        action={config.IMPORT_URL}
                        className='container'
                        accept='.xls,.xlsx'
                        drag
                        beforeUpload={this.handleBeforeUpload}
                        onSuccess={this.handleUploadSuccess}
                        onError={this.handleUploadError}
                        onChange={this.handleUploadChange}
                    >
                        <i className='el-icon-upload'/>
                        <div className='el-upload__text'>拖放文件，或 <em>点此上传</em></div>
                        <div className='el-upload__tip'>只能上传 xlsx/xls 文件；导入的数据没有进行去重处理</div>
                    </Upload>
                </Loading>
            </main>
        )
    }
}

export default Import
