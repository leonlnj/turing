"""
    Turing Minimal Openapi Spec for SDK

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from turing.generated.api_client import ApiClient, Endpoint as _Endpoint
from turing.generated.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from turing.generated.model.build_ensembler_image_request import BuildEnsemblerImageRequest
from turing.generated.model.ensembler_image_runner_type import EnsemblerImageRunnerType
from turing.generated.model.ensembler_images import EnsemblerImages


class EnsemblerImagesApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

        def __create_ensembler_image(
            self,
            project_id,
            ensembler_id,
            build_ensembler_image_request,
            **kwargs
        ):
            """Creates a new ensembler image  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.create_ensembler_image(project_id, ensembler_id, build_ensembler_image_request, async_req=True)
            >>> result = thread.get()

            Args:
                project_id (int):
                ensembler_id (int):
                build_ensembler_image_request (BuildEnsemblerImageRequest): A JSON object containing information about the ensembler

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                None
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['project_id'] = \
                project_id
            kwargs['ensembler_id'] = \
                ensembler_id
            kwargs['build_ensembler_image_request'] = \
                build_ensembler_image_request
            return self.call_with_http_info(**kwargs)

        self.create_ensembler_image = _Endpoint(
            settings={
                'response_type': None,
                'auth': [],
                'endpoint_path': '/projects/{project_id}/ensemblers/{ensembler_id}/images',
                'operation_id': 'create_ensembler_image',
                'http_method': 'PUT',
                'servers': None,
            },
            params_map={
                'all': [
                    'project_id',
                    'ensembler_id',
                    'build_ensembler_image_request',
                ],
                'required': [
                    'project_id',
                    'ensembler_id',
                    'build_ensembler_image_request',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'project_id':
                        (int,),
                    'ensembler_id':
                        (int,),
                    'build_ensembler_image_request':
                        (BuildEnsemblerImageRequest,),
                },
                'attribute_map': {
                    'project_id': 'project_id',
                    'ensembler_id': 'ensembler_id',
                },
                'location_map': {
                    'project_id': 'path',
                    'ensembler_id': 'path',
                    'build_ensembler_image_request': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client,
            callable=__create_ensembler_image
        )

        def __list_ensembler_images(
            self,
            project_id,
            ensembler_id,
            **kwargs
        ):
            """Returns a list of ensembler images that belong to the ensembler  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.list_ensembler_images(project_id, ensembler_id, async_req=True)
            >>> result = thread.get()

            Args:
                project_id (int):
                ensembler_id (int):

            Keyword Args:
                runner_type (EnsemblerImageRunnerType): [optional]
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                EnsemblerImages
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['project_id'] = \
                project_id
            kwargs['ensembler_id'] = \
                ensembler_id
            return self.call_with_http_info(**kwargs)

        self.list_ensembler_images = _Endpoint(
            settings={
                'response_type': (EnsemblerImages,),
                'auth': [],
                'endpoint_path': '/projects/{project_id}/ensemblers/{ensembler_id}/images',
                'operation_id': 'list_ensembler_images',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'project_id',
                    'ensembler_id',
                    'runner_type',
                ],
                'required': [
                    'project_id',
                    'ensembler_id',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'project_id':
                        (int,),
                    'ensembler_id':
                        (int,),
                    'runner_type':
                        (EnsemblerImageRunnerType,),
                },
                'attribute_map': {
                    'project_id': 'project_id',
                    'ensembler_id': 'ensembler_id',
                    'runner_type': 'runner_type',
                },
                'location_map': {
                    'project_id': 'path',
                    'ensembler_id': 'path',
                    'runner_type': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__list_ensembler_images
        )
