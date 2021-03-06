/* 
 * OpenAPI spec version: 2.4.9
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

type DomainAvailableError struct {

	// Short identifier for the error, suitable for indicating the specific error within client code
	Code string `json:"code,omitempty"`

	// Domain name
	Domain string `json:"domain,omitempty"`

	// Human-readable, English description of the error
	Message string `json:"message,omitempty"`

	// <ul> <li style='margin-left: 12px;'>JSONPath referring to a field containing an error</li> <strong style='margin-left: 12px;'>OR</strong> <li style='margin-left: 12px;'>JSONPath referring to a field that refers to an object containing an error, with more detail in `pathRelated`</li> </ul>
	Path string `json:"path,omitempty"`

	// HTTP status code that would return for a single check
	Status int32 `json:"status,omitempty"`
}
