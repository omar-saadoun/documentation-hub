-define(AEVM_PRIMOP_ERR_REASON_OOG(Resource, Gas, State), {out_of_gas, {{primop, Resource}, Gas}, State}).