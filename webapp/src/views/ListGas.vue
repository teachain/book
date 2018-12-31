<template>
    <div>
        <Card v-if="records.length==0">
             当前还没有挖掘机的加油记录
        </Card>
        <Card v-for="item in records" :value="item.id" :key="item.id" style="margin-top: 10px;margin-left: 5px;margin-right: 5px">
            <Form>
                <FormItem label="挖掘机" :label-width="80">
                    {{item.device_name}}
                </FormItem>
                <FormItem label="油量" :label-width="80">
                    {{item.amount}}升
                </FormItem>
                <FormItem label="油价" :label-width="80">
                    {{item.price}}升/元
                </FormItem>
                <FormItem label="总价格" :label-width="80">
                    {{item.total_price}}元
                </FormItem>
                <FormItem label="日期时间" :label-width="80">
                    {{item.create_time | dateFormat}}
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
        this.$http.get("/gas/list").then((data) => {
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