package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type enumValidator struct {
	allowedValues []string
}

func (v enumValidator) Description(_ context.Context) string {
	return "Only accepts the following values: " + strings.Join(v.allowedValues, ",")
}

func (v enumValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v enumValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	val, ok := req.AttributeConfig.(types.String)
	if !ok {
		resp.Diagnostics.AddError("attribute config is not a string", "")
		return
	}
	if val.Null {
		return
	}

	for _, candidate := range v.allowedValues {
		if candidate == val.Value {
			return
		}
	}
	resp.Diagnostics.AddError(fmt.Sprintf("invalid value %s. Value must be one of: %s", val.Value, strings.Join(v.allowedValues, ",")), "")
}
