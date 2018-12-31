<template>
    <div>
        <Card v-if="records.length==0">
             当前还没有挖掘机的出勤记录
        </Card>
        <Card v-for="item in records" :value="item.id" :key="item.id" style="margin-top: 10px;margin-left: 5px;margin-right: 5px">
            <Form>
                <FormItem label="挖掘机" :label-width="80">
                    {{item.device_name}}
                </FormItem>
                <FormItem label="司机" :label-width="80">
                    {{item.driver_name}}
                </FormItem>
                <FormItem label="客户" :label-width="80">
                    {{item.customer}}
                </FormItem>
                <FormItem label="地点" :label-width="80">
                    {{item.workplace}}
                </FormItem>
                <FormItem label="备注" :label-width="80">
                    {{item.remarks}}
                </FormItem>
                <FormItem label="开始时间" :label-width="80">
                    {{item.start_time | dateFormat}}
                </FormItem>
                <FormItem label="结束时间" :label-width="80">
                    {{item.end_time | dateFormat}}
                </FormItem>
            </Form>
        </Card>
        <Button type="primary" v-on:click="goBack()" style="margin-top: 10px">返回</Button>
    </div>
</template>
<script>
import moment from "moment"
export default {
    data() {
        return {
            records: [],
        }
    },
    filters: {
        dateFormat(val) {
            return moment(val).format('YYYY-MM-DD HH:mm');
        }
    },
    created() {
        this.$http.get("/record/list").then((data) => {
            this.records = data.data.data;
        });
    },
    methods: {
        goBack() {
            this.$router.push({ path: "/" });
        }
    }

}
</script>