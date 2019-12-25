<template>
  <el-container>
    <system-sidebar></system-sidebar>
    <el-main>
      <el-pagination
        background
        layout="prev, pager, next, sizes, total"
        :total="logTotal"
        :page-size="20"
        @size-change="changePageSize"
        @current-change="changePage"
        @prev-click="changePage"
        @next-click="changePage">
      </el-pagination>
      <el-table
        :data="logs"
        border
        ref="table"
        style="width: 100%">
        <el-table-column
          prop="id"
          label="ID">
        </el-table-column>
        <el-table-column
          prop="username"
          label="用户名">
        </el-table-column>
        <el-table-column
          prop="ip"
          label="登录IP">
        </el-table-column>
        <el-table-column
          prop="task_id"
          label="任务ID">
        </el-table-column>
        <el-table-column
          prop="task_name"
          label="任务名">
        </el-table-column>
        <el-table-column
          prop="action_name"
          label="动作类型">
        </el-table-column>
        <el-table-column
          prop="action"
          label="修改内容">
        </el-table-column>
        <el-table-column
          label="操作时间"
          width="">
          <template slot-scope="scope">
            {{scope.row.created | formatTime}}
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
import systemSidebar from './sidebar'
import systemService from '../../api/system'
export default {
  name: 'action-log',
  data () {
    return {
      logs: [],
      logTotal: 0,
      searchParams: {
        page_size: 20,
        page: 1
      }
    }
  },
  created () {
    this.search()
  },
  components: {systemSidebar},
  methods: {
    changePage (page) {
      this.searchParams.page = page
      this.search()
    },
    changePageSize (pageSize) {
      this.searchParams.page_size = pageSize
      this.search()
    },
    search () {
      systemService.actionLogList(this.searchParams, (data) => {
        this.logs = data.data
        this.logTotal = data.total
      })
    }
  }
}
</script>
